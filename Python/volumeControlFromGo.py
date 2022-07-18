from subps import set_master_volume, get_master_volume
import sys

args = sys.argv
first_parameter = args[1]

print(first_parameter)

if first_parameter == "get":
    get_master_volume()

else:
    set_master_volume(int(first_parameter))

