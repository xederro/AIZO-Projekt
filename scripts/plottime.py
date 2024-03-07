import csv
import os
import sys
import matplotlib.pyplot as plt
import numpy as np
from matplotlib.ticker import FuncFormatter
from scipy.interpolate import make_interp_spline


def read_csv(file_path):
    data = {}
    with open(file_path, 'r') as file:
        reader = csv.reader(file, delimiter=';')
        for row in reader:
            name, size, time_in_nanoseconds = row
            size = int(size)
            time_in_nanoseconds = int(time_in_nanoseconds)
            if name not in data:
                data[name] = {'size': [], 'time_in_nanoseconds': []}
            data[name]['size'].append(size)
            data[name]['time_in_nanoseconds'].append(time_in_nanoseconds)
    return data

def format_time_ticks(x, pos):
    units = ['ns', 'μs', 'ms', 's', 'min', 'hr', 'day']
    conversions = [1, 1000, 1000, 1000, 60, 60, 24]

    unit_index = 0
    while x >= 1000 and unit_index < len(units) - 1:
        x /= 1000
        unit_index += 1

    return f"{x:.2f} {units[unit_index]}"


def plot_data(data, save_folder):
    for name, values in data.items():
        plt.figure()
        x = values['size']
        y = values['time_in_nanoseconds']

        X_Y_Spline = make_interp_spline(x, y)
        X_ = np.linspace(min(x), max(x), 500)
        Y_ = X_Y_Spline(X_)
        plt.plot(X_, Y_, '-', color='blue')
        plt.scatter(x, y, color='blue')
        plt.title(name)
        plt.xlabel('Size')
        plt.gca().yaxis.set_major_formatter(FuncFormatter(format_time_ticks))
        save_path = os.path.join(save_folder, f"{name}.png")
        plt.savefig(save_path)
        plt.close()

def main(file_path, save_folder):
    os.makedirs(save_folder, exist_ok=True)
    data = read_csv(file_path)
    plot_data(data, save_folder)

if __name__ == "__main__":
    if (args_count := len(sys.argv)) > 3:
        print(f"One argument expected, got {args_count - 1}")
        raise SystemExit(2)
    elif args_count < 3:
        print("You must specify source and the target directory")
        raise SystemExit(2)
    main(sys.argv[1], sys.argv[2])
