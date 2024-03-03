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
    duration_s = x / 1e9

    if duration_s < 1e-3:
        return f"{x} ns"
    elif duration_s < 1e-3:
        return f"{x / 1e3:.2f} µs"
    elif duration_s < 1:
        return f"{duration_s * 1e3:.2f} ms"
    elif duration_s < 60:
        return f"{duration_s:.2f} s"
    elif duration_s < 3600:
        return f"{duration_s / 60:.2f} m"
    else:
        return f"{duration_s / 3600:.2f} h"


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
