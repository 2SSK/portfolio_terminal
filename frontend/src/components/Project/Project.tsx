import Link from "./Link";

import collegePreview from "/images/college-preview.png";
import coflowPreview from "/images/coflow-preview.png";

export default function Project() {
  return (
    <div className="mt-4 flex flex-col gap-4">
      <Link
        image={coflowPreview}
        github=""
        href="https://co-flow-sauravsinghkarmwars-projects.vercel.app/"
        title="Coflow"
        description="A collaborative platform empowering teams to ideate, plan, and innovate seamlessly on a virtual canvas."
      />
      <Link
        image={collegePreview}
        github=""
        href="https://react-college-website.vercel.app/"
        title="CollegePortal"
        description="Built a dynamic college website with multiple pages using ReactJS and TailwindCSS."
      />
    </div>
  );
}
