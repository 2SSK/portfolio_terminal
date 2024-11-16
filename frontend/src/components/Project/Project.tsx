import Link from "./Link";

export default function Project() {
  return (
    <div className="gap-10 md:gap-2">
      <Link
        href="https://photo-ai-ssk.vercel.app"
        title="PhotoAI"
        description="An AI-powered web app for seamless image and video editing with advanced tools like smart cropping, recoloring, and transcription."
      />
      <Link
        href="https://co-flow-sauravsinghkarmwars-projects.vercel.app/"
        title="Coflow"
        description="A collaborative platform empowering teams to ideate, plan, and innovate seamlessly on a virtual canvas."
      />
      <Link
        href="https://react-college-website.vercel.app/"
        title="CollegePortal"
        description="A responsive university website built with React.js"
      />
      <Link
        href="https://gemini-clone-ssk.vercel.app/"
        title="GeminiAI"
        description="A clone of GeminiAI website built with React.js"
      />
      <Link
        href="https://github.com/2SSK/dot-files/"
        title="i3wm Dotfiles"
        description="A complete setup for i3wm with configurations for Alacritty, Kitty, Neovim, Emacs, Polybar, Rofi, and more."
      />
      <Link
        href="https://github.com/2SSK/go-todo-cli-app"
        title="Go Todo CLI"
        description="A simple CLI app to manage todos built with Go"
      />
    </div>
  );
}
