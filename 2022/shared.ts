import fs from 'fs';
import path from 'path';

export const readFile = (filePath: string): string => {
    try {
        const fullPath = path.resolve(filePath);
        const data = fs.readFileSync(fullPath, 'utf8');
        return data;
    } catch (err) {
        console.error(`Error reading file from path: ${filePath}`);
        throw err;
    }
};