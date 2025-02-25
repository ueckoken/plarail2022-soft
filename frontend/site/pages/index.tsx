import type { NextPage } from "next"
import Head from "next/head"
import styles from "../styles/Home.module.css"
import RailroadMap from "../components/RailRoadMap"
import VideoCast from "../components/VideoCast"
import { useEffect, useRef, useState } from "react"
import {
  BlockId,
  blockIdMap,
  blockIds,
  BlockMessage,
  BlockStateIdMap,
  bunkiRailId,
  BunkiRailId,
  pointIdMap,
  pointIdMapReverse,
  StationId,
  StationMessage,
  StationState,
  StationStateIdMapReverse,
  StopRailId,
  stopRailId,
} from "../types/control-messages"

// OFF: false, ON: trueと対応
type StopPointState = Record<StopRailId, boolean>
const INITIAL_STOP_POINT_STATE: StopPointState = {
  shinjuku_s1: false,
  shinjuku_s2: false,
  sakurajosui_s0: false,
  sakurajosui_s1: false,
  sakurajosui_s2: false,
  sakurajosui_s3: false,
  sakurajosui_s4: false,
  sakurajosui_s5: false,
  chofu_s0: false,
  chofu_s1: false,
  chofu_s2: false,
  chofu_s3: false,
  chofu_s4: false,
  hachioji_s1: false,
  hachioji_s2: false,
  hashimoto_s1: false,
  hashimoto_s2: false,
}

type BlockState = Record<BlockId, boolean>
const INITIAL_BLOCK_STATE: BlockState = {
  shinjuku_b1: false,
  shinjuku_b2: false,
  sakurajosui_b1: false,
  sakurajosui_b2: false,
  sakurajosui_b3: false,
  sakurajosui_b4: false,
  sakurajosui_b5: false,
  sakurajosui_b6: false,
  chofu_b1: false,
  chofu_b2: false,
  chofu_b3: false,
  chofu_b4: false,
  chofu_b5: false,
  hashimoto_b1: false,
  hashimoto_b2: false,
  hachioji_b1: false,
  hachioji_b2: false,
  unknown: false,
}

type SwitchPointState = Record<BunkiRailId, boolean>
const INITIAL_SWITCH_POINT_STATE: SwitchPointState = {
  chofu_p1: false,
  sakurajosui_p1: false,
}

const Home: NextPage = () => {
  const stationWs = useRef<WebSocket>()
  const [stopPointState, setStopPointState] = useState<StopPointState>(
    INITIAL_STOP_POINT_STATE
  )
  const [switchPointState, setSwitchPointState] = useState<SwitchPointState>(
    INITIAL_SWITCH_POINT_STATE
  )
  useEffect(() => {
    setSwitchPointState(INITIAL_SWITCH_POINT_STATE)
  })
  const [blockState, setBlockState] = useState<BlockState>(INITIAL_BLOCK_STATE)
  const [selectedStationId, setSelectedStationId] =
    useState<StationId>("unknown")
  const [trainPosition1, setTrainPosition1] = useState<number>(0.4)

  const [roomIds, setRoomIds] = useState<string[]>(["chofu", "train"])

  const changeStopPointOrSwtichPointState = (
    stationId: StationId,
    state: StationState
  ) => {
    const message: StationMessage = {
      station: {
        stationId: pointIdMapReverse[stationId] as keyof typeof pointIdMap,
      },
      state: StationStateIdMapReverse[state] as keyof typeof BlockStateIdMap,
    }
    console.log(JSON.stringify(message))
    stationWs.current?.send(JSON.stringify(message))
  }
  const toggleStopPointOrSwitchPointState = (stationId: StationId) => {
    let state
    if (stopRailId.is(stationId)) {
      state = stopPointState[stationId]
    } else if (bunkiRailId.is(stationId)) {
      state = switchPointState[stationId]
    } else {
      return
    }
    const nextState = state ? "OFF" : "ON"
    changeStopPointOrSwtichPointState(stationId, nextState)
  }

  useEffect(() => {
    const ws = new WebSocket("wss://control.chofufes2022.ueckoken.club/pointws")
    stationWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
      console.log(e)
    })
    ws.addEventListener("message", (e) => {
      // console.log("recieved message")
      // console.log(e)
      const message: StationMessage = JSON.parse(e.data)
      // console.log(message.station.stationId)
      if (message.station.stationId === 0 || message.state === 0) {
        return
      }
      if (stopRailId.is(pointIdMap[message.station.stationId])) {
        setStopPointState((previousStopPointState) => ({
          ...previousStopPointState,
          [pointIdMap[message.station.stationId]]: message.state === 1,
        }))
        return
      }
      if (bunkiRailId.is(pointIdMap[message.station.stationId])) {
        setSwitchPointState((previousSwitchPointState) => ({
          ...previousSwitchPointState,
          [pointIdMap[message.station.stationId]]: message.state === 1,
        }))
      }
    })
    ws.addEventListener("error", (e) => {
      console.log("error occured")
      console.log(e)
    })
    ws.addEventListener("close", (e) => {
      console.log("closed")
      console.log(e)
    })
    return () => {
      ws.close()
    }
  }, [])

  useEffect(() => {
    const ws = new WebSocket("wss://control.chofufes2022.ueckoken.club/blockws")
    stationWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
      console.log(e)
    })
    ws.addEventListener("message", (e) => {
      // console.log("recieved message")
      // console.log(e)
      const message: BlockMessage = JSON.parse(e.data)
      // console.log(message)
      if (message.blockId === 0 || message.state === 0) {
        return
      }
      if (blockIds.is(blockIdMap[message.blockId])) {
        setBlockState((previousStopPointState) => ({
          ...previousStopPointState,
          [blockIdMap[message.blockId]]: message.state === 2,
        }))
        return
      }
    })
    ws.addEventListener("error", (e) => {
      console.log("error occured")
      console.log(e)
    })
    ws.addEventListener("close", (e) => {
      console.log("closed")
      console.log(e)
    })
    return () => {
      ws.close()
    }
  }, [])

  useEffect(() => {
    const intervalId = setInterval(() => {
      const tmpTrainPosition1 = trainPosition1 + 0.01
      setTrainPosition1(tmpTrainPosition1 <= 1 ? tmpTrainPosition1 : 0)
    }, 20)
    return () => clearInterval(intervalId)
  })

  return (
    <div className={styles.container}>
      <Head>
        <title>工研&times;鉄研プラレール展示 管理画面</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/kokenLogo.ico" />
      </Head>

      <header>
        <h1 className={styles.title}>工研&times;鉄研プラレール展示 管理画面</h1>
      </header>

      <main className={styles.main}>
        <section>
          <h2>映像部分</h2>
          <div
            style={{
              margin: 0,
              padding: 0,
              width: "100%",
              position: "relative",
            }}
          >
            <VideoCast
              roomIds={roomIds}
              styles={[
                {
                  position: "relative",
                  zIndex: 1,
                  width: "100%",
                },
                {
                  position: "absolute",
                  zIndex: 2,
                  left: 0,
                  width: "33%",
                },
                {
                  position: "absolute",
                  zIndex: 2,
                  left: "33%",
                  width: "33%",
                },
                {
                  position: "absolute",
                  zIndex: 2,
                  left: "66%",
                  width: "33%",
                },
                {
                  position: "absolute",
                  zIndex: 2,
                  top: "50%",
                  left: 0,
                  width: "33%",
                },
                {
                  position: "absolute",
                  zIndex: 2,
                  top: "50%",
                  left: "33%",
                  width: "33%",
                },
              ]}
            />
          </div>
          <button
            onClick={() => {
              setRoomIds(["hachioji", "train"])
            }}
          >
            京王八王子
          </button>
          <button
            onClick={() => {
              setRoomIds(["kitano", "train"])
            }}
          >
            橋本
          </button>
          <button
            onClick={() => {
              setRoomIds(["hashimoto", "train"])
            }}
          >
            調布
          </button>
          <button
            onClick={() => {
              setRoomIds(["chofu", "train"])
            }}
          >
            桜上水
          </button>
          <button
            onClick={() => {
              setRoomIds(["shinjyuku", "train"])
            }}
          >
            新宿
          </button>
          <button
            onClick={() => {
              setRoomIds(["train"])
            }}
          >
            車両前景
          </button>
          <button
            onClick={() => {
              setRoomIds([
                "none",
                "hachioji",
                "chofu",
                "hashimoto",
                "sakurajosui",
                "shinjuku",
              ])
            }}
          >
            all
          </button>
        </section>

        <section>
          <h2>地図部分</h2>
          <RailroadMap
            datas={{
              stop: stopPointState,
              switchState: switchPointState,
              blockState: blockState,
              train1: {
                positionScale: trainPosition1,
                id: "koken",
              },
            }}
            onStopPointOrSwitchPointClick={(stationId) => {
              toggleStopPointOrSwitchPointState(stationId)
            }}
          />
        </section>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://www.koken.club.uec.ac.jp/"
          target="_blank"
          rel="noopener noreferrer"
        >
          &copy;2022 電気通信大学工学研究部
        </a>
      </footer>
    </div>
  )
}

export default Home
