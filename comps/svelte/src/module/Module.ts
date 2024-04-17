export interface Module {
    path: string;
    component: () => Promise<{ default: any }>;
}

export class App {
    modules: Module[];
}