class DataLayer(object):
    """ The purpose of the class is to be a data layer over the JSON data. """
    def get_athletes(data):
        """ Returns the JSON part with the athletes. """
        return data["athletes"]

    def get_squad_id(athlete):
        return athlete["squad_id"]

    def get_squad_name_by_id(data, id):
        for squad in data["squads"]:
            if squad["id"] == id:
                return squad["name"]

    def get_squads(data):
        """ Returns the JSON part with the squads. """
        return data["squads"]
