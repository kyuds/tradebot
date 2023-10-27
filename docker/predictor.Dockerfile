FROM pytorch/pytorch:2.1.0-cuda12.1-cudnn8-runtime

# Some Python settings to reduce load
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Install pip dependencies
ENV PIP_IGNORE_INSTALLED 0
COPY ./predictor .
RUN pip install --upgrade pip
RUN pip install --user --no-cache-dir -r ./dependencies/requirements.txt

# Will be ran in docker-compose
# CMD ["python", "predictor.py"]
