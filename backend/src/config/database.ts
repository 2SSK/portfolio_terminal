import mongoose from "mongoose";
import dotenv from "dotenv";

dotenv.config();

const mongoURI: any = process.env.MONGO_URI;

const connectDB = async () => {
  try {
    await mongoose.connect(mongoURI as string);
    console.log("MongoDB connected successfully");
  } catch (error: any) {
    console.error("Error connecting MongoDB", error.message);
    process.exit(1);
  }
};

export default connectDB;
