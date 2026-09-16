FROM python:3.14-slim-trixie
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY app.py .
EXPOSE 32777
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "32777"]
