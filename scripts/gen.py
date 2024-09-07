import pandas as pd
from faker import Faker

def create_csv_with_names(num_rows, filename):
    fake = Faker()

    names = [fake.name() for _ in range(num_rows)]

    df = pd.DataFrame(names, columns=["Name"])

    df.to_csv(filename, index=False)


nums = [100, 1000, 10000, 100000]

for num in nums:
    create_csv_with_names(num, f"./data/names_{num}.csv")
