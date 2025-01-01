import { useState } from "react";
import { User } from "../../types/user";
import { Plus, Pencil, Trash2 } from "lucide-react";

interface UserTableProps {
  users: User[];
  onEdit: (user: User) => void;
  onDelete: (user: User) => void;
  onNew: () => void;
}

export default function UserTable({
  users,
  onEdit,
  onDelete,
  onNew,
}: UserTableProps) {
  const [selectedUser, setSelectedUser] = useState<User | null>(null);

  return (
    <div className="p-6">
      <div className="mb-4 flex justify-between items-center">
        <h1 className="text-3xl font-extrabold text-gray-900">
          User Management
        </h1>
        <div className="flex space-x-2">
          <button
            onClick={onNew}
            className="bg-blue-500 text-white px-4 py-2 rounded flex items-center gap-2 hover:bg-blue-600"
          >
            <Plus size={20} /> New
          </button>
          <button
            onClick={() => selectedUser && onEdit(selectedUser)}
            disabled={!selectedUser}
            className="bg-blue-500 text-white px-4 py-2 rounded flex items-center gap-2 hover:bg-blue-600 disabled:bg-gray-300"
          >
            <Pencil size={20} /> Edit
          </button>
          <button
            onClick={() => selectedUser && onDelete(selectedUser)}
            disabled={!selectedUser}
            className="bg-blue-500 text-white px-4 py-2 rounded flex items-center gap-2 hover:bg-blue-600 disabled:bg-gray-300"
          >
            <Trash2 size={20} /> Delete
          </button>
        </div>
      </div>

      <div className="bg-white rounded-lg shadow">
        <table className="min-w-full">
          <thead>
            <tr className="border-b bg-gray-200">
              <th className="px-6 py-3 text-left text-gray-700">ID</th>
              <th className="px-6 py-3 text-left text-gray-700">Username</th>
              <th className="px-6 py-3 text-left text-gray-700">Age</th>
              <th className="px-6 py-3 text-left text-gray-700">Created At</th>
              <th className="px-6 py-3 text-left text-gray-700">Updated At</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => (
              <tr
                key={user.ID}
                onClick={() => setSelectedUser(user)}
                className={`border-b cursor-pointer hover:bg-gray-100 ${
                  selectedUser?.ID === user.ID ? "bg-blue-100" : "bg-white"
                }`}
              >
                <td className="px-6 py-4 text-gray-900">{user.ID}</td>
                <td className="px-6 py-4 text-gray-900">{user.Username}</td>
                <td className="px-6 py-4 text-gray-900">{user.Age}</td>
                <td className="px-6 py-4 text-gray-900">
                  {new Date(user.CreatedAt).toLocaleDateString()}
                </td>
                <td className="px-6 py-4 text-gray-900">
                  {new Date(user.UpdatedAt).toLocaleDateString()}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
