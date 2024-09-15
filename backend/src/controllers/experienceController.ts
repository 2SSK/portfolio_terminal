import { Request, Response } from "express";
import Experience from "../models/Experience";

export const createExperience = async (req: Request, res: Response) => {
  try {
    const experience = new Experience(req.body);
    await experience.save();
    res.status(201).json(experience);
  } catch (error) {
    res.status(400).json({ error: "Unable to create experience" });
  }
};

export const getExperiences = async (req: Request, res: Response) => {
  const experiences = await Experience.find();
  res.json(experiences);
};
