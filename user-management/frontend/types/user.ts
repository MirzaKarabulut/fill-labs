export interface User {
  ID: number;
  Username: string;
  Age: number;
  CreatedAt: string;
  UpdatedAt: string;
}

export type FormMode = 'create' | 'edit' | 'delete';