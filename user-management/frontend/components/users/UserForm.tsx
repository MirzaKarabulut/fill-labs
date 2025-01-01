import { useState } from "react";
import { User, FormMode } from "../../types/user";
import { ChevronLeft } from "lucide-react";

interface UserFormProps {
  user?: User;
  mode: FormMode;
  onSubmit: (user: Partial<User>) => void;
  onBack: () => void;
}

export default function UserForm({
  user,
  mode,
  onSubmit,
  onBack,
}: UserFormProps) {
  const [formData, setFormData] = useState<Partial<User>>({
    Username: user?.Username || "",
    Age: user?.Age || 0,
  });

  const buttonText = {
    create: "Create",
    edit: "Save",
    delete: "Delete",
  }[mode];

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!/^[a-zA-Z]+$/.test(formData.Username || "")) {
      alert("Username must contain only letters.");
      return;
    }
    if (formData.Age === undefined || formData.Age < 0 || formData.Age > 100) {
      alert("Age must be between 0 and 100.");
      return;
    }
    onSubmit({
      Username: formData.Username,
      Age:
        typeof formData.Age === "string"
          ? parseInt(formData.Age)
          : formData.Age,
    });
  };

  return (
    <div className="p-6">
      <div className="mb-4">
        <button
          onClick={onBack}
          className="text-blue-500 flex items-center gap-2 hover:text-blue-600"
        >
          <ChevronLeft size={20} /> Back to List
        </button>
      </div>

      <div className="bg-white rounded-lg shadow p-6 max-w-md mx-auto">
        <h2 className="text-2xl font-bold text-gray-900 mb-4">
          {mode.charAt(0).toUpperCase() + mode.slice(1)} User
        </h2>

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Username
            </label>
            <input
              type="text"
              value={formData.Username || ""}
              onChange={(e) =>
                setFormData({ ...formData, Username: e.target.value })
              }
              disabled={mode === "delete"}
              className="w-full px-3 py-2 border rounded text-gray-900 font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Age
            </label>
            <input
              type="number"
              value={formData.Age || ""}
              onChange={(e) =>
                setFormData({
                  ...formData,
                  Age: e.target.value ? parseInt(e.target.value) : 0,
                })
              }
              disabled={mode === "delete"}
              className="w-full px-3 py-2 border rounded text-gray-900 font-medium focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>

          <div className="flex justify-end space-x-2">
            <button
              type="button"
              onClick={onBack}
              className="px-4 py-2 border rounded hover:bg-gray-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              className={`px-4 py-2 rounded text-white ${
                mode === "delete"
                  ? "bg-red-500 hover:bg-red-600"
                  : "bg-blue-500 hover:bg-blue-600"
              }`}
            >
              {buttonText}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
