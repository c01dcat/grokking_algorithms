states_needed = {"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

stations = {}
stations["kone"] = {"id", "nv", "ut"}
stations["ktwo"] = {"wa", "id", "mt"}
stations["kthree"] = {"or", "nv", "ca"}
stations["kfour"] = {"nv", "ut"}
stations["kfive"] = {"ca", "az"}

final_stations = set()

if __name__ == '__main__':
    while states_needed:
        best_station = None
        states_covered = set()
        for station, states in stations.items():
            covered = states_needed.intersection(states)  # states_needed & states
            if len(covered) > len(states_covered):
                best_station = station
                states_covered = covered

        states_needed -= states_covered
        final_stations.add(best_station)

    print(final_stations)
