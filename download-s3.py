import os
import boto3
import tkinter as tk
from tkinter import filedialog

aws_access_key_id = 'key_id'
aws_secret_access_key = 'secret_key'

# Create an S3, I already configured the CLI with the credentials
s3 = boto3.client('s3', aws_access_key_id=aws_access_key_id, aws_secret_access_key=aws_secret_access_key)

folder_name = "uploaded"
# List objects in the bucket
response = s3.list_objects_v2(Bucket = 'final-project-test-bucket', Prefix=f"{folder_name}/")

path = filedialog.askdirectory(initialdir=os.getcwd(), title="Select a dir", )
file_list = [obj['Key'] for obj in response.get('Contents', [])][1:]


# Download each file in the bucket
for file in file_list:
    local_file_path = os.path.join(path, file[9:])
    # Download the file
    s3.download_file('final-project-test-bucket', file, local_file_path)

