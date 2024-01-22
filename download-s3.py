import os
import boto3
import tkinter as tk
from tkinter import filedialog


# Create an S3, I already configured the CLI with the credentials
s3 = boto3.client('s3')

# List objects in the bucket
response = s3.list_objects_v2(Bucket = 'final-project-test-bucket')

path = filedialog.askdirectory(initialdir=os.getcwd(), title="Select a dir", )

# Download each file in the bucket
for obj in response.get('Contents', []):
    file_key = obj['Key']
    local_file_path = os.path.join(path, file_key)
    # Download the file
    s3.download_file('final-project-test-bucket', file_key, local_file_path)

