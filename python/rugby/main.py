import json

from constants import STD_TEAM_CRITERIA
from application_layer import ApplicationLayer
from presentation_layer import production


PATH = "rugby_athletes.json_bck"


def read_from_hdd(path):
    with open(path) as f:
        return json.load(f)


def main():
    engine = ApplicationLayer(read_from_hdd(PATH), STD_TEAM_CRITERIA)
    for (squad_id, squad_name, team) in engine.generate_teams():
        production(squad_id, squad_name, team)

if __name__ == "__main__":
    main()
