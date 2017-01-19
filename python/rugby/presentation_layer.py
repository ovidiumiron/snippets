def production(squad_id, squad_name, team):
    print("squad_id: {0}, squad name: {1}".format(squad_id, squad_name))
    print("------------------------------")
    for member in team:
        print("position: {0}, name: {1}".format(
            member["position"], member["name"]))
    print("------------------------------")
