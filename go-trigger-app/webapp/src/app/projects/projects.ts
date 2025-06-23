export interface Project {
  name: string;
  github?: string;
  website?: string;
  image: string;
  npm?: string;
}

export const projects: Project[] = [
  {
    name: 'Project1',
    github: 'https://github.com/Izaacapp/my-awesome-project',
    website: 'https://myawesomeproject.com',
    image: 'assets/projects/my-awesome-image.png',
    npm: ''
  },
  {
    name: 'Project2',
    github: 'https://github.com/YourUsername/project2',
    website: 'https://project2website.com',
    image: 'assets/projects/project2-image.png',
    npm: 'https://www.npmjs.com/package/project2'
  }
];
