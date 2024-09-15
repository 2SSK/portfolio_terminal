import { Request, Response } from "express";
import path from "path";

export const downloadResume = (req: Request, res: Response) => {
  const filePath = path.join(__dirname, "../../uploads/resume.pdf");
  res.download(filePath);
};
