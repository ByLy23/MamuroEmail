export interface EmailInterface {
  id: string;
  time: string;
  content: string;
  directory: string;
  fileName: string;
  user: string;
}
export interface contentInterface {
  hits: number;
  content: EmailInterface[];
}
