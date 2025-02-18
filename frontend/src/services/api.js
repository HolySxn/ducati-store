export const API_URL = 'http://localhost:8080';

export const fetchData = async (endpoint) => {
  const response = await fetch(`${API_URL}/${endpoint}`);
  if (!response.ok) {
    throw new Error('Network response was not ok');
  }
  return response.json();
};

export const deleteItem = async (endpoint, id) => {
  const response = await fetch(`${API_URL}/${endpoint}/${id}`, {
    method: 'DELETE',
  });
  if (!response.ok) {
    throw new Error('Failed to delete item');
  }
  return response.json();
};

export const createItem = async (endpoint, data) => {
  const response = await fetch(`${API_URL}/${endpoint}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    throw new Error('Failed to create item');
  }
  return response.json();
};

export const updateItem = async (endpoint, id, data) => {
  const response = await fetch(`${API_URL}/${endpoint}/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  });
  if (!response.ok) {
    throw new Error('Failed to update item');
  }
  return response.json();
};