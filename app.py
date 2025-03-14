from fastapi import FastAPI
import psycopg2
from psycopg2 import extras
from contextlib import asynccontextmanager
import logging
from psycopg2 import pool
import random
# from config import config

# ....CODE....
uvicorn_error = logging.getLogger("uvicorn.error")
uvicorn_error.disabled = True
uvicorn_access = logging.getLogger("uvicorn.access")
uvicorn_access.disabled = True

from contextlib import asynccontextmanager

app = FastAPI()

# Create a connection pool at the start of the application
db_pool = None

def init_db_pool():
    global db_pool
    try:
        # Initialize the connection pool
        print('Initializing the PostgreSQL connection pool...')
        db_pool = psycopg2.pool.SimpleConnectionPool(
            1,  # minconn (minimum connections in the pool)
            10, # maxconn (maximum connections in the pool)
            dbname="postgres",
            user="postgres",
            password="root",
            host="localhost"
        )
        if db_pool:
            print("Connection pool created successfully.")
    except Exception as e:
        print(f"Error initializing the connection pool: {e}")

@asynccontextmanager
async def lifespan(app: FastAPI):
    # Initialization at the startup of the application
    init_db_pool()
    yield
    # Optional cleanup at shutdown
    if db_pool:
        print("Closing the connection pool.")
        db_pool.closeall()

app = FastAPI(lifespan=lifespan)

def connect():
    """ Connect to the PostgreSQL database server using the pool """
    conn = None
    db_version = None
    try:
        # Get a connection from the pool
        conn = db_pool.getconn()
        if conn:
            # create a cursor
            cur = conn.cursor()
            # execute a statement
            cur.execute('SELECT version()')
            # display the PostgreSQL database server version
            db_version = cur.fetchone()
            cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(f"Error: {error}")
    finally:
        if conn:
            # Return the connection to the pool
            db_pool.putconn(conn)
    return db_version

@app.get("/python-db")
async def get_pg_db_version():
    return connect()

@app.get("/python-rand")
async def rand():
    return {"number": random.randint(1, 10)}

@app.get("/python-health")
async def health():
    return "Hello World"