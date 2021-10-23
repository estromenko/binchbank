import pandas as pd

from faker import Faker
from random import randint

faker = Faker()

types = ['credit', 'deposit', 'lising', 'etc']

clients = []
operations = []
employees = []
branches = []

for i in range(1, 4):
    branches.append([
        i,
        faker.name(),
    ])

for i in range(1, 51):
    clients.append([
        i,
        faker.name(),
        faker.date(),
        faker.boolean()
    ])

for i in range(1, 21):
    operations.append([
        i,
        randint(1, 51),
        randint(1, 10000),
        types[randint(0, 3)],
        faker.date()
    ])

for i in range(1, 11):
    employees.append([
        i,
        faker.name(),
        randint(1, 3),
        randint(1, 3),
    ])

# Client
pd.DataFrame(clients).to_csv('./clients.csv', header=[
    'id',
    'username',
    'date',
    'is_legal_entity'
], index=False)

# Operations
pd.DataFrame(operations).to_csv('./operations.csv', header=[
    'id',
    'client_id',
    'amount',
    'type',
    'date'
], index=False)

# Employees
pd.DataFrame(employees).to_csv('./employees.csv', header=[
    'id',
    'username',
    'branch_id',
    'plan'
], index=False)

# Branches
pd.DataFrame(branches).to_csv('./branches.csv', header=[
    'id',
    'manager_name'
], index=False)