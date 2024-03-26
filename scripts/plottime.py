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
            algorithm, data_type, array, modifier, size, time_in_nanoseconds = row
            size = int(size)
            time_in_nanoseconds = int(time_in_nanoseconds)
            name = f"{algorithm} {data_type} {array}"
            if modifier != '':
                name += f" {modifier}"
            if name not in data:
                data[name] = {'data': [algorithm, data_type, array, modifier], 'size': [], 'time_in_nanoseconds': []}
            data[name]['size'].append(size)
            data[name]['time_in_nanoseconds'].append(time_in_nanoseconds)
    return data

def plot_specific_data(data, save_folder):
    # Define the combinations of data entries to plot
    specific_plots = [
        # ShellSort
        {'algorithm': 'ShellSort', 'data_type': 'int8', 'modifier': 'Lazarus', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'int8', 'modifier': 'Shell', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'int32', 'modifier': 'Lazarus', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'int32', 'modifier': 'Shell', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'int64', 'modifier': 'Lazarus', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'int64', 'modifier': 'Shell', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'float32', 'modifier': 'Lazarus', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'float32', 'modifier': 'Shell', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'float64', 'modifier': 'Lazarus', 'array': None},
        {'algorithm': 'ShellSort', 'data_type': 'float64', 'modifier': 'Shell', 'array': None},
        # InsertionSort
        {'algorithm': 'InsertionSort', 'data_type': 'int8', 'modifier': None, 'array': None},
        {'algorithm': 'InsertionSort', 'data_type': 'int32', 'modifier': None, 'array': None},
        {'algorithm': 'InsertionSort', 'data_type': 'int64', 'modifier': None, 'array': None},
        {'algorithm': 'InsertionSort', 'data_type': 'float64', 'modifier': None, 'array': None},
        {'algorithm': 'InsertionSort', 'data_type': 'float32', 'modifier': None, 'array': None},
        # HeapSort
        {'algorithm': 'HeapSort', 'data_type': 'int8', 'modifier': None, 'array': None},
        {'algorithm': 'HeapSort', 'data_type': 'int32', 'modifier': None, 'array': None},
        {'algorithm': 'HeapSort', 'data_type': 'int64', 'modifier': None, 'array': None},
        {'algorithm': 'HeapSort', 'data_type': 'float64', 'modifier': None, 'array': None},
        {'algorithm': 'HeapSort', 'data_type': 'float32', 'modifier': None, 'array': None},
        # QuickSort
        {'algorithm': 'QuickSort', 'data_type': 'int8', 'modifier': 'First Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int8', 'modifier': 'Last Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int8', 'modifier': 'Random Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int8', 'modifier': 'Middle Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int32', 'modifier': 'First Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int32', 'modifier': 'Last Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int32', 'modifier': 'Random Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int32', 'modifier': 'Middle Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int64', 'modifier': 'First Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int64', 'modifier': 'Last Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int64', 'modifier': 'Random Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'int64', 'modifier': 'Middle Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float32', 'modifier': 'First Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float32', 'modifier': 'Last Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float32', 'modifier': 'Random Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float32', 'modifier': 'Middle Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float64', 'modifier': 'First Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float64', 'modifier': 'Last Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float64', 'modifier': 'Random Pivot', 'array': None},
        {'algorithm': 'QuickSort', 'data_type': 'float64', 'modifier': 'Middle Pivot', 'array': None},
    ]

    for plot_spec in specific_plots:
        plt.figure(figsize=(10, 6))
        for _, values in data.items():
            algorithm, data_type, array, modifier = values['data']
            if (plot_spec['algorithm'] is None or plot_spec['algorithm'] == algorithm) and \
               (plot_spec['data_type'] is None or plot_spec['data_type'] == data_type) and \
               (plot_spec['modifier'] is None or plot_spec['modifier'] == modifier) and \
               (plot_spec['array'] is None or plot_spec['array'] == array):
                x = values['size']
                y = values['time_in_nanoseconds']
                X_Y_Spline = make_interp_spline(x, y)
                X_ = np.linspace(min(x), max(x), 500)
                Y_ = X_Y_Spline(X_)
                label = ""
                if plot_spec['algorithm'] is None:
                    label += f"{algorithm} "
                if plot_spec['data_type'] is None:
                    label += f"{data_type} "
                if plot_spec['array'] is None:
                    label += f"{array} "
                if plot_spec['modifier'] is None:
                    label += f"{modifier}"
                plt.plot(X_, Y_, '-', label=label)
                plt.scatter(x, y)

        label = ""
        if plot_spec['algorithm'] is not None:
            label += f"{plot_spec['algorithm']} "
        if plot_spec['data_type'] is not None:
            label += f"{plot_spec['data_type']} "
        if plot_spec['array'] is not None:
            label += f"{plot_spec['array']} "
        if plot_spec['modifier'] is not None:
            label += f"{plot_spec['modifier']} "

        label = label.strip()
#         plt.title(label)
        plt.xlabel('Size')
        plt.ylabel('Time (nanoseconds)')
        plt.legend()
        plt.gca().yaxis.set_major_formatter(FuncFormatter(format_time_ticks))
        save_path = os.path.join(save_folder, f"{label}.png")
        plt.savefig(save_path)
        plt.close()

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
        plt.figure(figsize=(10, 6))
        x = values['size']
        y = values['time_in_nanoseconds']

        X_Y_Spline = make_interp_spline(x, y)
        X_ = np.linspace(min(x), max(x), 500)
        Y_ = X_Y_Spline(X_)
        plt.plot(X_, Y_, '-', color='blue')
        plt.scatter(x, y, color='blue')
#         plt.title(name)
        plt.xlabel('Size')
        plt.gca().yaxis.set_major_formatter(FuncFormatter(format_time_ticks))
        save_path = os.path.join(save_folder, f"{name}.png")
        plt.savefig(save_path)
        plt.close()


def main(file_path, save_folder):
    os.makedirs(save_folder, exist_ok=True)
    data = read_csv(file_path)
    plot_data(data, save_folder)
    plot_specific_data(data, save_folder)

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python script.py <source_csv_file> <target_directory>")
        sys.exit(1)
    main(sys.argv[1], sys.argv[2])
