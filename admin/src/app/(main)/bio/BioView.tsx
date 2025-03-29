import { BioResponse } from "@/actions/bio_action";
import { BioBox } from "@/components/BioBox";
import Form from "next/form"; // Assuming this is a custom or future Next.js utility

type BioViewProps = {
  bio: BioResponse | null;
  error: string | undefined;
  onUpdateBio: (formData: FormData) => Promise<void>;
};

export default function BioView({ bio, error, onUpdateBio }: BioViewProps) {
  async function handleSubmit(formData: FormData) {
    "use server";
    if (!formData.get("image") && bio?.image) {
      formData.set("image", bio.image);
    }
    await onUpdateBio(formData);
  }

  return (
    <div className="container mx-auto p-6 pt-20 flex flex-col gap-8 md:flex-row md:items-start">
      {/* Bio Display */}
      <BioBox bio={bio} />

      {/* Bio Update Form */}
      <Form
        action={handleSubmit}
        className="flex-1 bg-white shadow-lg rounded-lg p-6 space-y-6"
      >
        <h3 className="text-lg font-semibold text-gray-700">Update Bio</h3>

        {error && (
          <div className="rounded-lg bg-red-50 p-3 text-sm text-red-600">
            {error}
          </div>
        )}

        <div className="space-y-1">
          <label
            htmlFor="image"
            className="block text-sm font-medium text-gray-700"
          >
            Profile Image
          </label>
          <input
            id="image"
            type="file"
            name="image"
            accept="image/*"
            className="block w-full border rounded text-sm text-gray-500 file:mr-4 file:rounded-md file:border-0 file:bg-blue-50 file:px-4 file:py-2 file:text-blue-700 hover:file:bg-blue-100"
          />
        </div>

        <div className="space-y-1">
          <label
            htmlFor="title"
            className="block text-sm font-medium text-gray-700"
          >
            Title <span className="text-red-500">*</span>
          </label>
          <input
            id="title"
            type="text"
            name="title"
            defaultValue={bio?.title ?? ""}
            required
            className="block w-full border rounded-md border-gray-300 p-2 focus:border-blue-500  focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div className="space-y-1">
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-700"
          >
            Name
          </label>
          <input
            id="name"
            type="text"
            name="name"
            defaultValue={bio?.name ?? ""}
            className="block w-full border rounded-md border-gray-300 p-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div className="space-y-1">
          <label
            htmlFor="description"
            className="block text-sm font-medium text-gray-700"
          >
            Description
          </label>
          <textarea
            id="description"
            name="description"
            defaultValue={bio?.description ?? ""}
            rows={3}
            className="block w-full border rounded-md border-gray-300 p-2 focus:border-blue-500 focus:ring-2 focus:ring-blue-500 resize-none"
          />
        </div>

        <div className="text-center">
          <button
            type="submit"
            className="rounded-md bg-blue-600 px-4 py-2 text-white transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
          >
            Update Bio
          </button>
        </div>
      </Form>
    </div>
  );
}
