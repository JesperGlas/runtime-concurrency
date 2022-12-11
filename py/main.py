from typing import Tuple
from math import sin, cos, sqrt
from multiprocessing import Pool
import random
import time

JOBS = 10 # How many jobs each version will schedule
BATCHES = 10 # How large each batch will be for concurrency

def calculate(id: str, run_time_seconds: float=1) -> None:
    """Function that simulates a heavy computation

    Args:
        id (str): An id for the task.
            Ex. Sequential (1/10)
                Concurrent (4/5)
        run_time_seconds (float): Simulated runtime of
            function in seconds
    """
    print(f"Starting {id}..")
    time.sleep(run_time_seconds)
    print(f"{id} done..")
    
def map_calculate(item: Tuple) -> None:
    """Help function to run calculate with an index in a map function

    Args:
        item (Tuple): Resulting index, value tuple from enumerate
    """
    i, _ = item
    calculate(f"Concurrent {i+1}/{JOBS}")

def sequential(total: int) -> None:
    """Sequential run of calculate n times

    Args:
        total (int): Total amount of jobs
    """
    for i in range(total):
        calculate(f"Sequential ({i+1}/{total})")

def concurrent(total: int) -> None:
    """Concurrent run of calculate n times

    Args:
        total (int): Defines how many times calculate will be called
    """
    result = [None for _ in range(total)]
    with Pool(BATCHES) as pool:
        pool.map(map_calculate, enumerate(result))

if __name__ == "__main__":
    print("Python Runtimes:")
    
    # Sequential timed
    start = time.perf_counter()
    sequential(JOBS)
    duration = time.perf_counter() - start
    print(f"Sequential: {duration} s")

    # Concurrent timed
    start = time.perf_counter()
    concurrent(JOBS)
    duration = time.perf_counter() - start
    print(f"Concurrent: {duration} s")
