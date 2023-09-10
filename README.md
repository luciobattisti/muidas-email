# muidas-email

## Run gomail program

### Build [MacOs]

```bash
go build -o gomail.so
```
### Run

```bash
./gomail.so -html-template=gotemplates/snowfall-release.html -newsletter=../data/sample-emails.csv -subject='Hello from Muidas' -from=<your-email> -token=<your-email-token> 
```

## Run email validator

### Create environment and install dependencies

```bash
conda create -n email-env -y python=3.9.1 jupyter ipykernel
conda activate email-env
pip install pandas
```

### Run 

python validate_email.py