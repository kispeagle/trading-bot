from sys import getsizeof
import websocket

data = []

def on_message(ws, message):
    print(message)

def on_error(ws, error):
    print(error)

def on_close(ws, close_status_code, close_msg):
    print("### closed ###")

def on_open(ws):
    print("Opened connection")

ws = websocket.WebSocketApp("wss://stream.binance.com:9443/ws/btcusdt@aggTrade",
                              on_open=on_open,
                              on_message=on_message,
                              on_error=on_error,
                              on_close=on_close)
ws.run_forever()  # Set dispatcher to automatic reconnection
# rel.signal(2, rel.abort)  # Keyboard Interrupt
# rel.dispatch()