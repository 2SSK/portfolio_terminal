import Link from "./Link";

import coflowPreview from "/images/projects/coflow-preview.png";
import quickCartPreview from "/images/projects/quickCart-preview.png";
import collegePreview from "/images/projects/college-preview.png";
import textEditorPreview from "/images/projects/texteditor-preview.png";
import containerPreview from "/images/projects/container_preview.png";

export default function Project() {
  return (
    <div className="mt-4 ml-10 flex flex-col gap-4">
      <Link
        image={coflowPreview}
        github="https://github.com/2SSK/CoFlow.git"
        href="https://co-flow-sauravsinghkarmwars-projects.vercel.app/"
        title="Coflow"
        description="A collaborative platform empowering teams to ideate, plan, and innovate seamlessly on a virtual canvas."
      />
      <Link
        image={quickCartPreview}
        github="https://github.com/2SSK/quick-cart.git"
        href="https://quick-cart-ssk.vercel.app/"
        title="Quick Cart (wip)"
        description="Built a seamless e-commerce platform for buyers and sellers, featuring secure authentication with Clerk, efficient user data sync via Inngest, and optimized product image management using Cloudinary. Powered by MongoDB, it ensures a smooth and scalable shopping experience. "
      />
      <Link
        image={collegePreview}
        github="https://github.com/2SSK/react-college-website.git"
        href="https://react-college-website.vercel.app/"
        title="CollegePortal"
        description="Built a dynamic college website with multiple pages using ReactJS and TailwindCSS."
      />
      <Link
        image={textEditorPreview}
        github="https://github.com/2SSK/text_editor.git"
        href="https://github.com/2SSK/text_editor.git"
        title="TextEditor GUI"
        description="Built a lightweight and functional text editor built using C++ and the FLTK library. This project demonstrates clean, modular design and effective use of modern C++ features for creating GUI-based applications."
      />
      <Link
        image={containerPreview}
        github="https://github.com/2SSK/container_from_scratch.git"
        href="https://github.com/2SSK/container_from_scratch.git"
        title="Container from Scratch"
        description="Built a simple container in Go form scratch. It includes a basic setup for running an isolated process within a minimal filesystem."
      />
    </div>
  );
}
