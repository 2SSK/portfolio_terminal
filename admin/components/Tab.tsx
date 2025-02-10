import React from "react";

const Tab = ({ setTab }: { setTab: (tab: string) => void }) => {
  return (
    <div className="w-full mt-4">
      <div className="flex justify-center">
        <ul className="flex gap-8 text-lg">
          <Link setTab={setTab} tabName={"Bio"} />
          <Link setTab={setTab} tabName={"Socials"} />
          <Link setTab={setTab} tabName={"Contacts"} />
          <Link setTab={setTab} tabName={"Resume"} />
          <Link setTab={setTab} tabName={"Projects"} />
          <Link setTab={setTab} tabName={"Experience"} />
        </ul>
      </div>
    </div>
  );
};

const Link = ({
  setTab,
  tabName,
}: {
  setTab: (tab: string) => void;
  tabName: string;
}) => {
  return (
    <li className="tab-link" onClick={() => setTab(tabName)}>
      {tabName}
    </li>
  );
};

export default Tab;
