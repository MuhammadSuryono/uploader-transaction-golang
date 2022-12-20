import pandas as pd

def csv_to_json(csv_path, json_path):
    df = pd.read_csv(csv_path, sep=';')
    df.to_json(json_path, orient='records')

if __name__ == '__main__':
    csv_path = r'C:\Users\Arg\Downloads\sales_data\SALES_19122022.csv'
    json_path = 'SALES_19122022.json'
    csv_to_json(csv_path, json_path)