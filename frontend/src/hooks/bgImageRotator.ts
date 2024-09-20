import { useEffect, useMemo } from "react";
import { useSetRecoilState } from "recoil";
import { bgState } from "../store/atom/atom";

import image1 from "/images/wallpapers/image1.jpg";
import image2 from "/images/wallpapers/image2.png";
import image3 from "/images/wallpapers/image3.png";
import image4 from "/images/wallpapers/image4.jpg";
import image5 from "/images/wallpapers/image5.png";

const BackgroundImageRotator = () => {
  const setBgImage = useSetRecoilState(bgState);

  // Memoize the image array to avoid recreating it on each render
  const images = useMemo(() => [image1, image2, image3, image4, image5], []);

  useEffect(() => {
    const changeBackgroundImage = () => {
      const randomImage = images[Math.floor(Math.random() * images.length)];
      setBgImage(randomImage);
    };

    // Initial image setting
    changeBackgroundImage();

    // Set an interval to change the background image every minute (60 seconds)
    const interval = setInterval(changeBackgroundImage, 1000 * 60);

    // Cleanup interval on component unmount
    return () => clearInterval(interval);
  }, [images, setBgImage]);

  return null;
};

export default BackgroundImageRotator;
