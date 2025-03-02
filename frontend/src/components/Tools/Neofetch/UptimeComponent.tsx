import { useEffect, useState } from "react";

const UptimeComponent = () => {
  const [uptime, setUptime] = useState("");

  useEffect(() => {
    const startDate = new Date("2003-06-15");
    const updateUptime = () => {
      const now = new Date();
      const differenceInTime = now.getTime() - startDate.getTime();
      const differenceInDays = differenceInTime / (1000 * 3600 * 24);
      const years = Math.floor(differenceInDays / 365);
      const remainingDays = Math.floor(differenceInDays % 365);

      setUptime(`${years} years, ${remainingDays} days`);
    };

    updateUptime();
    const intervalId = setInterval(updateUptime, 1000 * 60 * 60 * 24);

    return () => clearInterval(intervalId);
  }, []);

  return <>{uptime}</>;
};

export default UptimeComponent;
