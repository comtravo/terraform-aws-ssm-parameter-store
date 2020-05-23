FROM comtravo/terraform:py3-0.12.25-1.0.0

WORKDIR /opt/terraform
RUN pip install awscli-local

COPY . .
