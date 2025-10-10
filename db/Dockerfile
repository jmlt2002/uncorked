FROM postgres:16-alpine

# Set environment variables (can be overridden by docker-compose or at runtime)
ENV POSTGRES_DB=uncorked
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

# Copy the initialization script

COPY init.sql /docker-entrypoint-initdb.d/
RUN chown -R postgres:postgres /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5432
