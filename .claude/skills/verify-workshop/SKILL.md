---
name: verify-workshop
description: Run a thorough, from-scratch end-to-end verification of the Kubernetes Introduction Workshop. Tears down and recreates the `workshop` kind cluster, walks through every module's every documented command against a live cluster, compares real output to documented expected output, and produces per-module drift reports under `verification/`. Use when the user asks to verify the workshop, check it works, test modules, or validate docs against reality.
---

# verify-workshop

Authoritative, non-skippable verification of the Kubernetes Introduction Workshop in this repository. The deliverable is a set of drift reports — one per module plus a summary — that distinguish real behavioral drift (wrong status, broken command, missing field) from cosmetic output differences (random pod suffixes, ClusterIPs, timestamps).

The target audience of the workshop is **Kubernetes beginners**. That shapes the verification: a "cosmetic" difference that would confuse a noob (e.g., a wildly different event order, an unexpected warning line, a prompt that says something different from the doc) is NOT cosmetic for this audience — flag it.

## When to use

Trigger on any of:
- "verify the workshop", "test the workshop", "check the workshop works"
- "run the full verification", "re-run verification", "verify from scratch"
- "does module N still work", "check module N against docs"
- explicit invocation via `/verify-workshop` or similar

If the user only asks about one module, follow the per-module flow below but skip the cluster rebuild and prior-module commands unless they ask for "from scratch."

## Non-negotiable rules

1. **Thorough means thorough.** Every fenced `bash` block in every targeted module must be executed. No skipping "because it's similar to the last one." If the user says "don't skip any step," do not skip any step.
2. **Run commands one block at a time and compare to the doc.** After each block, diff real stdout/stderr against the following "Expected output" block (or inline expected phrase). Record matches and mismatches.
3. **Distinguish behavioral drift from cosmetic drift.** Pod random suffixes, ClusterIPs, random NodePorts, ephemeral ports, timestamps, age columns, event ordering, and minor formatting are *cosmetic*. Different resource kinds reported, missing fields, commands that fail outright, counts that don't match (`3/3 ready` vs `0/3 ready`), wrong YAML shape — those are *behavioral*.
4. **Beginner-lens check.** For each step, also ask: "would a Kubernetes newcomer who copy-pastes this be confused?" Warning lines, deprecated-API notices, or Helm NOTES that differ from the doc count as real drift for this audience even if technically cosmetic. Note these under **Beginner-confusion drift**.
5. **Write drift reports as you go.** One `verification/drift-module-NN.md` per module, appended after that module finishes. Do not batch to the end — partial progress must survive a session interruption.
6. **Modules build on each other's state.** Do NOT reset the cluster between modules. The README module order 1 → 11 is load-bearing (the Ingress module deploys the four services that Helm later reinstalls; the configuration module mutates the `web-app` deployment, etc.). Cleaning up resources between modules is the workshop's job, not yours.
7. **This is a local kind cluster.** You may freely apply manifests, exec into pods, delete namespaces, restart deployments, helm upgrade Traefik. You may NOT push to remotes, mutate shared infra, or touch anything outside this repo / local Docker.
8. **No placeholder outputs.** If a command fails unexpectedly or output doesn't match, record the real output verbatim. Never fabricate expected-looking output.
9. **Scope discipline.** Do not fix drifts during a verification run — only record. Fixing is a separate pass the user can request after reading the reports. Do not run **Challenge Exercise** blocks; they are explicitly student-driven and not part of the verified path. Do run Host-based routing and the optional OOMKilled demo — the README lists them as part of the module flow, just marked optional.

## Required context to load first

Read, in this order:
1. `README.md` — workshop structure, 11 modules, expected outputs, inline notes
2. `manifests/00-kind-config.yaml` — cluster topology (single control-plane, ports 30080/30443)
3. `manifests/*.yaml` — all workshop manifests (resolve any "apply this file" instruction to its contents before executing so you know what the apply should do)
4. `helm/example-app/` — chart structure and default values
5. Recent `git log` — note any recent fixes (a fresh run must still pass; stale drifts may have been patched)

## Workflow

### Phase 0 — Environment check

Confirm required tools are installed. Hard-fail if any are missing:
```
for t in docker kind kubectl helm curl jq; do command -v "$t" >/dev/null || { echo "MISSING: $t"; exit 1; }; done
docker info >/dev/null 2>&1 || { echo "Docker daemon unreachable (is Colima/Docker Desktop running?)"; exit 1; }
```

Record tool versions (`kind version`, `kubectl version --client`, `helm version`, `docker version --format '{{.Server.Version}}'`).

### Phase 1 — Cluster rebuild (from-scratch runs only)

```
kind delete cluster --name workshop 2>/dev/null || true
kind create cluster --name workshop --config manifests/00-kind-config.yaml
kubectl cluster-info
kubectl get nodes
```

Confirm one control-plane node `Ready`, kubectl context is `kind-workshop`.

Also run the README's "Install Metrics Server" step here — it's a prerequisite used from Module 8 onward:
```
helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
helm repo update
helm install metrics-server metrics-server/metrics-server \
  -n kube-system \
  --set 'args[0]=--kubelet-insecure-tls' \
  --set 'args[1]=--metric-resolution=10s' \
  --set 'args[2]=--kubelet-request-timeout=5s'
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=metrics-server -n kube-system --timeout=300s
kubectl top nodes
```
If `kubectl top nodes` fails on the first try, wait ~60s and retry — that's the documented behavior.

### Phase 2 — Sprint plan

Write `.claude/sprint-criteria.md` with done criteria (every module drift report exists, SUMMARY exists, cluster in expected end-state) and max iteration count. Create one TodoWrite task per module + setup + summary. Start with Module 1.

### Phase 3 — Per-module loop

For each module 1–11 in `README.md`:

1. **Read the module end-to-end first.** Note:
   - Commands that mutate cluster state (`apply`, `delete`, `scale`, `set image`, `rollout`, `edit`, `helm install/upgrade/rollback/uninstall`) vs read-only (`get`, `describe`, `logs`, `top`).
   - "Expected output" phrases — these are often inline, not fenced (`"You should see a pod named web-app with status Running"`).
   - Cross-module dependencies (Module 4 deletes the standalone pod; Module 5 now deletes the standalone RS before applying the Deployment; Module 6 installs Traefik; Module 9 mutates the `web-app` Deployment; Module 11 installs releases that coexist with manifest-deployed resources).
   - Interactive / blocking steps: `kubectl exec -ti`, `kubectl port-forward`, `kubectl edit`, `watch`, `kubectl logs -f`. Substitute with a non-interactive equivalent and note it in the drift report:
     - `exec -ti … -- ash` → `exec -- sh -c '<fixed command>'`
     - `port-forward` → run in background with `&`, curl, then `kill`
     - `edit cm/...` → `kubectl patch cm ...` with an equivalent change
     - `watch kubectl top pod` → a single `kubectl top pod`
     - `kubectl logs -f` → `kubectl logs --tail=20`

2. **Execute exercises in order.** Run each fenced `bash` block via the `Bash` tool. For long operations use explicit `--timeout`:
   - Rollouts: `kubectl rollout status --timeout=180s`
   - Pod readiness: `kubectl wait --for=condition=ready pod -l … --timeout=300s`
   - Helm installs: pass `--wait --timeout 5m` when the doc doesn't already

3. **Inspect output immediately.** Compare real output to the doc's expectation:
   - **Match** → move on.
   - **Cosmetic mismatch** (random suffix, different IP, different age, different node port in the 30000-32767 range) → move on, record under Cosmetic.
   - **Behavioral mismatch** (wrong count, wrong status, command fails, wrong field) → record in drift report with exact README line reference, command run, expected text, actual text, hypothesis, suggested fix.
   - **Beginner-confusion mismatch** (deprecation warning, unexpected Helm NOTES text, command output shape that doesn't match the screenshots/examples in README) → record under Beginner-confusion drift.

4. **Module-specific things to check** (seen on prior runs of similar Kind-based workshops):
   - Module 2 (Pods): `kubectl exec -ti web-app -- ash` — confirm the image still ships an Alpine shell. If not, the doc note claiming "our images are based on Alpine" is drift.
   - Module 3 (Services): confirm the NodePort is in the 30000-32767 range and NOT 30080/30443 (only those two are host-mapped, per the recently-added Kind note). Confirm port-forward still works.
   - Module 4 (ReplicaSets): `kubectl get rs` should show `web-app` with `DESIRED=3 CURRENT=3 READY=3`. After deleting the static pod, pod count must stay at 3.
   - Module 5 (Deployments): the `kubectl delete rs web-app` step must actually delete a ReplicaSet (it was added in a recent fix — if missing in docs, that's drift). `kubectl get rs` after Deployment apply should show a single `web-app-<hash>` ReplicaSet, not two. Rolling updates should show the old RS scaled to 0 and a new RS at 3. `kubectl rollout undo` must return the previous image.
   - Module 6 (Ingress): Traefik install via Helm — check the labels (`app.kubernetes.io/name=traefik`) still match; newer Traefik charts sometimes change defaults. Confirm `curl http://localhost:30080/` returns 200 (the multi-app ingress reserves `/` for `web-app`). Skip the `/etc/hosts` edit if running non-interactively — document that it was skipped.
   - Module 7 (Logging): `--prefix` label output, `--since`, `--previous`. For `--previous`, you must first cause a crash; if none happened, log that the command produced the "previous terminated container not found" error.
   - Module 8 (Resources): after `kubectl apply -f manifests/06-deployment-with-resources.yaml`, `kubectl get pod -l app=web-app -o yaml | grep -A 6 resources:` must show `50Mi` requests and `100Mi` limits (post-fix values). The optional OOMKilled demo should actually produce `OOMKilled` in status — if it doesn't, that's behavioral drift worth flagging for the beginner audience.
   - Module 9 (Configuration): confirm all 3 ConfigMaps and the Secret are created. `kubectl exec … env | grep -E 'OTHER_NAME|NAME|ENV_VAR'` must show values. `cat /tmp/config/file.txt` must return `I'm in a file`.
   - Module 10 (Headlamp): desktop app; verify only the kubectl-visible side (no launcher needed). Skip GUI steps and note them as skipped.
   - Module 11 (Helm): `helm dependency update ./helm/example-app` must succeed against the valkey repo. Check that `helm install example-app-ruby` reports `STATUS: deployed`. Verify `helm rollback example-app-ruby 1` drops the ingress. Install Python/Go/Rust releases, then `helm uninstall` for each at the end.

5. **Write `verification/drift-module-NN.md`** with these sections:
   - `Status: PASS | PASS with N minor doc drifts | FAIL`
   - `## Executed` — terse bullet per exercise (command + one-line outcome)
   - `## Behavioral drift (real)` — numbered, each with:
     - One-line description
     - Exact command and README location (`README.md:<line>`)
     - Expected output (quoted from README)
     - Actual output (quoted, truncated if huge)
     - Root cause hypothesis
     - Suggested fix with file:line reference
   - `## Beginner-confusion drift` — same shape as behavioral, but flagged as "technically cosmetic, confusing for noobs"
   - `## Cosmetic differences (not drift)` — bulleted
   - `## Suggested fixes` — consolidated list, if any

6. **Mark the module task completed** and move to the next.

### Phase 4 — Final summary

Write `verification/SUMMARY.md`:
- Date, cluster image (`kindest/node:vX.Y.Z`), kubectl/helm/kind versions
- Overall result (PASS / PASS with drifts / FAIL)
- Per-module status table with drift counts (behavioral, beginner-confusion, cosmetic)
- Top-N drifts ranked by beginner-reader impact (not by technical severity)
- Cluster end-state observations: running deployments, ingress rules, helm releases, namespace resource counts
- One-paragraph recommendation

### Phase 5 — Cleanup

Delete `.claude/sprint-criteria.md`. Leave the cluster running (the user may want to poke at it). Ask: "Would you like to commit the verification artifacts?" Do not auto-commit.

## Baseline expectations (update after each clean run)

At end-state after a clean full run:
- 1 Kind cluster `workshop`, 1 control-plane node, Ready
- `workshop` namespace set as current context
- Deployments: `web-app` (3 replicas, configured via ConfigMaps+Secret), `web-app-ruby`, `web-app-python`, `web-app-go`, `web-app-rust` (2 replicas each)
- Services: same 5 + whatever Helm created during Module 11
- Ingress: `web-app` — final state depends on the last applied manifest (`05-ingress.yaml`, `05-ingress-multipath.yaml`, or `05-ingress-hostbased.yaml`). Record which.
- 3 ConfigMaps + 1 Secret from `07-configuration.yaml`, plus `kube-root-ca.crt`
- Helm releases in `workshop` namespace (Module 11): depends on whether student ran uninstall at the end
- `kube-system`: `metrics-server`, `traefik`
- All 11 manifest files apply cleanly with `kubectl apply --dry-run=server -f manifests/` (skip `00-kind-config.yaml` — it's a Kind config, not a K8s resource)

If a new run produces drift beyond this baseline, flag it prominently in SUMMARY.md.

## Deliverable

After a successful run:
- `verification/drift-module-01.md` through `drift-module-11.md` (11 files)
- `verification/SUMMARY.md`
- Cluster in expected end-state described above
