import csv

def text_to_txt(input_file, output_file, delimiter=' '):
    with open(input_file, 'r') as f_input, open(output_file, 'w', newline='') as f_output:
        writer = csv.writer(f_output)
        for line in f_input:
          writer.writerow(line.split())
            

# Example usage:
input_file = 'export-HSK1.txt'
output_file = 'HSK1.csv'
text_to_txt(input_file, output_file)
