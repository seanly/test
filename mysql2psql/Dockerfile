FROM python:2

RUN apt update -y && apt install git && \
    git clone -b gitlab https://github.com/gitlabhq/mysql-postgresql-converter.git /code

# https://gitlab.com/gitlab-org/gitlab/-/blob/v9.0.6/doc/update/mysql_to_postgresql.md