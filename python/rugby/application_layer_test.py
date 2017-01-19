import unittest

from application_layer import ApplicationLayer


class TestApplicationLayer(unittest.TestCase):
    def test_generate_teams(self):
        data = {"athletes": [
            {'position': 'prop', 'is_injured': False, "squad_id": 0},
            {'position': 'prop', 'is_injured': False, "squad_id": 3}
        ], "squads": [
            {"id": 0, "name": "a"},
            {"id": 3, "name": "b"},
        ]
        }

        team_criteria = {
            "position": {"prop": 1},
            "is_injured": {False: 1}
        }
        engine = ApplicationLayer(data, team_criteria)
        teams = list(engine.generate_teams())

        self.assertEqual(
            [(0, 'a',
              [{'position': 'prop', 'is_injured': False, 'squad_id': 0}]),
             (3, 'b',
              [{'position': 'prop', 'is_injured': False, 'squad_id': 3}])],
            teams)

    def test_generate_teams_no_members(self):
        data = {"athletes": [
            {'position': 'prop', 'is_injured': False, "squad_id": 0},
            {'position': 'prop', 'is_injured': False, "squad_id": 3}
        ], "squads": [
            {"id": 0, "name": "a"},
            {"id": 3, "name": "b"},
        ]
        }

        team_criteria = {
            "position": {"prop": 2},
            "is_injured": {False: 1}
        }
        engine = ApplicationLayer(data, team_criteria)
        teams = list(engine.generate_teams())

        self.assertEqual([], teams)

if __name__ == '__main__':
    unittest.main()
