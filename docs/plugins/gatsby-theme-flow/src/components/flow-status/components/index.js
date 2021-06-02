import React, { useEffect, useState } from "react";
import StatusCard from "./statusCard";

import { getNextSporkDate, getNetworkVersion } from "../helpers";

function NetworkVersion(networkName) {
  const [version, setVersion] = useState("");
  useEffect(() => {
    async function get() {
      try {
        const v = await getNetworkVersion(networkName);
        setVersion(v);
      } catch (e) {
        console.log("Error getting network version...");
      }
    }
    get();
  }, []);
  return (
    <>
      <h4>Version</h4>
      {version}
    </>
  );
}

function NextSporkDate(networkName) {
  const [nextSpork, setNextSpork] = useState("");
  useEffect(() => {
    async function get() {
      try {
        const d = await getNextSporkDate(networkName);
        setNextSpork(d);
      } catch (e) {
        console.log("Error getting next spork date...");
      }
    }
    get();
  }, []);
  return (
    <>
      <h4>Next Spork</h4>
      {nextSpork}
    </>
  );
}

export { StatusCard, NetworkVersion, NextSporkDate };
