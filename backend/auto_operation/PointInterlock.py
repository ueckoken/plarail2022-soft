from Components import Junction
from State import State


# 列車の通過中など、ポイントを切り替えてはいけないときに切り替わらないよう制御する
class PointInterlock:
    def __init__(self, state: State, TRAINLENGTH: float):
        self.__state = state
        self.__TRAINLENGTH = TRAINLENGTH

    # ポイントを切り替える
    def requestToggle(self, junctionId: Junction.JunctionId):
        junction = self.__state.getJunctionById(junctionId)
        trainOnJunction = self.__state.getTrainInSection(junction.getOutSection())
        if trainOnJunction is None or trainOnJunction.mileage > self.__TRAINLENGTH:
            junction.toggle()
            print(f"[PointInterlock.requestToggle] junction {junction.id} toggled!")
