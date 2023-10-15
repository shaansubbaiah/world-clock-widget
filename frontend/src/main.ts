import "./style.css";

import { GetClocks } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";

let clocks: main.Clock[] = [];

const loadConfig = async () => {
  clocks = await GetClocks();
  console.log(JSON.stringify(clocks));
};

const updateClock = () => {
  document.getElementById("clock-1")!.innerText = "";
  const now = new Date();
  clocks.forEach((c) => {
    let timestring = now.toLocaleTimeString("en-GB", {
      timeZone: c.timezone,
    });
    // console.log(timestring);
    document.getElementById("clock-1")!.innerText +=
      timestring + " " + c.name + "\n";
  });
};

loadConfig();
setInterval(updateClock, 500);
