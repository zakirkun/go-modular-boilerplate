#!/bin/bash

echo "Stopping and removing containers..."
docker-compose down

echo "Do you want to remove volumes as well? (This will delete all data) [y/N]"
read -r remove_volumes

if [[ "$remove_volumes" =~ ^[Yy]$ ]]; then
    echo "Removing volumes..."
    docker-compose down -v
    echo "Volumes removed."
fi

echo "Do you want to remove all related Docker images? [y/N]"
read -r remove_images

if [[ "$remove_images" =~ ^[Yy]$ ]]; then
    echo "Removing images..."
    docker rmi $(docker images -q backend_modules:latest mysql:8.0) 2>/dev/null || true
    echo "Images removed."
fi

echo "Cleanup completed."