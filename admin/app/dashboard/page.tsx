"use client";
import React, { useState } from "react";
import Tab from "@/components/Tab";
import Bio from "@/components/Bio";
import Socials from "@/components/Socials";
import Contacts from "@/components/Contacts";
import Resume from "@/components/Resume";
import Experience from "@/components/Experience";
import Projects from "@/components/Projects";

const Page = () => {
  const [tab, setTab] = useState("Bio");

  return (
    <>
      <Tab setTab={setTab} />
      {(tab === "Bio" && <Bio />) ||
        (tab === "Socials" && <Socials />) ||
        (tab === "Contacts" && <Contacts />) ||
        (tab === "Resume" && <Resume />) ||
        (tab === "Experience" && <Experience />) ||
        (tab === "Projects" && <Projects />)}
    </>
  );
};

export default Page;
