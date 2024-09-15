import { Request, Response } from "express";
import Project from "../models/Project";
import multer from "multer";

const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    cb(null, "uploads/");
  },
  filename: (req, file, cb) => {
    cb(null, `${Date.now()}-${file.originalname}`);
  },
});
const upload = multer({ storage });

export const createProject = [
  upload.single("thumbnail"),
  async (req: Request, res: Response) => {
    try {
      const { title, description, repoLink, hostLink } = req.body;
      const project = new Project({
        title,
        description,
        thumbnail: req.file?.path,
        repoLink,
        hostLink,
      });
      await project.save();
      res.status(201).json(project);
    } catch (error) {
      res.status(400).json({ error: "Unable to create project" });
    }
  },
];

export const getProjects = async (req: Request, res: Response) => {
  const projects = await Project.find();
  res.json(projects);
};
