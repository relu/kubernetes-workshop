FROM python:3.13-alpine

COPY requirements.txt /app/
WORKDIR /app
RUN pip install --no-cache-dir -r requirements.txt

COPY app.py /app/

EXPOSE 3000

CMD ["python", "app.py"]
