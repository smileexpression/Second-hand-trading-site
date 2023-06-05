export async function uploadImage(formData) {
  try {
    const response = await fetch('http://localhost:8080/images', {
      method: 'POST',
      body: formData,
    });
    // 如果响应状态码为 200 OK，表示上传成功
    if (response.ok) {
      const data = await response.json();
      return data;
    } else {
      throw new Error('上传失败');
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function getImage(id) {
  try {
    const imageResponse = await fetch(`http://localhost:8080/images/${id}`);
    if (imageResponse.ok) {
      const blob = await imageResponse.blob();
      const url = URL.createObjectURL(blob);
      return url;
    } else {
      throw new Error('获取图片失败');
    }
  } catch (error) {
    console.error(error);
    throw error;
  }
}

export async function convertImageToBlob(imageFile) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      const blob = new Blob([reader.result], { type: imageFile.type });
      resolve(blob);
    };
    reader.onerror = reject;
    reader.readAsArrayBuffer(imageFile);
  });
}