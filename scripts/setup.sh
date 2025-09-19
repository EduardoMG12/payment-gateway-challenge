#!/bin/bash

echo "Starting project setup..."

# Copying .env.example
cp .env.example .env

# Installing npm dependencies
npm install

# Initializing Husky
npx husky init

# set up pre-commit hook
echo "npx lint-staged" > .husky/pre-commit

echo "Setup complete!"
echo "Remember to run 'npx husky add .husky/pre-commit \"npm run pre-commit\"' after this."