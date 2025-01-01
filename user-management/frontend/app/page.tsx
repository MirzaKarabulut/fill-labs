"use client";

import { useState, useEffect } from "react";
import { User,FormMode } from "../types/user";
import UserTable from "@/components/users/UserTable";
import UserForm from "@/components/users/UserForm";
import { api } from "@/services/api";

export default function Home() {
  const [users, setUsers] = useState<User[]>([]);
  const [view, setView] = useState<"list" | "form">("list");
  const [formMode, setFormMode] = useState<FormMode>("create");
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchUsers = async () => {
    try {
      setLoading(true);
      const data = await api.getUsers();
      setUsers(data);
      setError(null);
    } catch (err) {
      setError("Failed to fetch users");
      console.error("Error fetching users:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleNew = () => {
    setFormMode("create");
    setSelectedUser(null);
    setView("form");
  };

  const handleEdit = (user: User) => {
    setFormMode("edit");
    setSelectedUser(user);
    setView("form");
  };

  const handleDelete = (user: User) => {
    setFormMode("delete");
    setSelectedUser(user);
    setView("form");
  };

  const handleSubmit = async (userData: Partial<User>) => {
    try {
      setLoading(true);

      if (formMode === "create") {
        await api.createUser(userData);
      } else if (formMode === "edit" && selectedUser) {
        await api.updateUser(selectedUser.ID, userData);
      } else if (formMode === "delete" && selectedUser) {
        await api.deleteUser(selectedUser.ID);
      }

      await fetchUsers();
      setView("list");
      setSelectedUser(null);
    } catch (err) {
      setError("Operation failed");
      console.error("Error submitting form:", err);
    } finally {
      setLoading(false);
    }
  };

  if (loading && users.length === 0) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-gray-500">Loading...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-red-500">{error}</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {view === "list" ? (
        <UserTable
          users={users}
          onEdit={handleEdit}
          onDelete={handleDelete}
          onNew={handleNew}
        />
      ) : (
        <UserForm
          user={selectedUser || undefined}
          mode={formMode}
          onSubmit={handleSubmit}
          onBack={() => {
            setView("list");
            setSelectedUser(null);
          }}
        />
      )}
    </div>
  );
}
