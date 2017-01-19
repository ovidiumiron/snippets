from collections import defaultdict
from copy import deepcopy
from data_layer import DataLayer as dl


def athlete_match(athlete, athlete_criteria):
    """
    Returns True if athlete matchs ALL athlete_criteria.

    >>> athlete = {"criteria1":"criteria11", "criteria2": True, "criteria3": 4}
    >>> athlete_criteria = {"criteria1": ["criteria11", "criteria12"], "criteria2": [True]}
    >>> athlete_match(athlete, athlete_criteria)
    True
    >>> athlete_criteria = {"criteria1": ["criteria11", "criteria12"], "criteria_not_in_athlete": [True]}
    >>> athlete_match(athlete, athlete_criteria)
    False
    """
    for key, value in athlete_criteria.items():
        if key not in athlete or athlete[key] not in value:
            return False
    return True


def make_athlete_criteria(team_criteria):
    """
    Starting from team_criteria makes the athlete_criteria.

    >>> team_criteria = {"criteria1": {"criteria11": 2, "criteria12":1}, "criteria2": {False: 11}}
    >>> make_athlete_criteria(team_criteria)
    defaultdict(<class 'list'>, {'criteria1': ['criteria11', 'criteria12'], 'criteria2': [False]})
    >>> team_criteria = {"criteria1": {"criteria11": 2, "criteria12":1}, "criteria2": {False: 11, True:4}}
    >>> make_athlete_criteria(team_criteria)
    defaultdict(<class 'list'>, {'criteria1': ['criteria11', 'criteria12'], 'criteria2': [False, True]})
    >>> team_criteria = {"criteria1": {"criteria11": 2, "criteria12":1}, "criteria2": {False: 0, True:4}}
    >>> make_athlete_criteria(team_criteria)
    defaultdict(<class 'list'>, {'criteria1': ['criteria11', 'criteria12'], 'criteria2': [False, True]})
    """
    athlete_criteria = defaultdict(list)

    for key, attributes in team_criteria.items():
        for attribute in attributes:
            athlete_criteria[key].append(attribute)
    return athlete_criteria


class ApplicationLayer(object):
    """
    The purpose of the class is to generate teams which respect team_criteria,
    the public function generate_teams.

    The algorithm is:
        - on the __init__ create self.data_engine which keeps a deepcopy of the
          team criteria
        - loop into the athletes and search for one which respect the
          team_criteria
        - decrease with 1 each criterion from team_criteria corresponding with
          the athlete( the function __update_team_criteria)
    """
    def __init__(self, data, team_criteria):
        """
        Create a data structure for the ApplicationLayer.
        "32": {
            "team_criteria": team_criteria,
            "team_members": []
            }
        where
            - 32 is the squad id
            .....
        """
        self.data = data
        self.athletes = dl.get_athletes(data)
        self.data_engine = dict()
        for squad in dl.get_squads(data):
            self.data_engine[squad["id"]] = {
                "team_criteria": deepcopy(team_criteria),
                "team_members": list()}

    def generate_teams(self):
        """ Yield teams satisfying criteria. """
        for ath in self.__extract_athletes():
            squad_id = dl.get_squad_id(ath)
            self.data_engine[squad_id]["team_members"].append(ath)

            """ We have already meet all criteria. """
            if self.__have_team(squad_id):
                yield (squad_id,
                       dl.get_squad_name_by_id(self.data, squad_id),
                       self.data_engine[squad_id]['team_members'])

    def __have_team(self, squad_id):
        """
        True is there is selected all the players for the team for squad_id.
        Else return False.
        """
        return not self.data_engine[squad_id]["team_criteria"]

    def __extract_athletes(self):
        """ Yield athletes which match team criteria. """
        for ath in self.athletes:
            if dl.get_squad_id(ath) not in self.data_engine:
                # Athlete has no squad. Just skip over it.
                continue

            team_criteria = \
                self.data_engine[dl.get_squad_id(ath)]["team_criteria"]

            if not team_criteria:
                # Probably already generated a team for athlete["squad_id"]
                continue

            if athlete_match(ath, make_athlete_criteria(team_criteria)):
                self.__update_team_criteria(team_criteria, ath)
                yield ath

    @staticmethod
    def __update_team_criteria(team_criteria, athlete):
        """
        Decrease with 1 all the criteria from team_criteria founded  in the
        athlete.
        Remove all the criteria which are empty.
        >>> team_criteria = {"criteria1": {"criteria11": 2, "criteria12":1}, "criteria2": {False: 11}}
        >>> athlete = {"criteria1": "criteria11",  "criteria2": False}
        >>> ApplicationLayer._ApplicationLayer__update_team_criteria(team_criteria, athlete)
        >>> team_criteria
        {'criteria1': {'criteria11': 1, 'criteria12': 1}, 'criteria2': {False: 10}}
        """
        for key, value in team_criteria.items():
            value[athlete[key]] -= 1
            if value[athlete[key]] == 0:
                del value[athlete[key]]

        # Remove keys with empty value.
        for k in list(team_criteria):
            if not team_criteria[k]:
                del team_criteria[k]


if __name__ == "__main__":
    import doctest
    doctest.testmod()
