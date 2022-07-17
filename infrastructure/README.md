# Infrastructure
Set up to work with AWS.

## Steps
Assuming you are in the root of the infrastructure folder.
Set up the remote states
1. ```cd remote-state```
2. ```terraform init```
3. ```terraform apply```

Set up all the environments
1. ```cd env-setup```
2. ```terraform init```
3. ```terraform apply```

Set up the application infrastructure
1. ```cd fghub/[dev|test|staging|prod]```
2. ```terraform init```
3. ```terraform apply```

For now application infrastructure, everything runs off of 1 folder, but it
may be separated later if it becomes too unwieldy.