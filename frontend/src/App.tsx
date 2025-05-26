import axios from "axios";
import { useState } from "react";
import { useEffect } from "react";
import Swal from "sweetalert2";
import qrcodeImage from "./assets/qrcode_example.png";

function App() {
  interface Product {
    product_id: string;
  }

  const BASE_URL = "http://localhost:8080/api/v1";
  const URL_ADD_PRODUCT = `${BASE_URL}/product`;
  const URL_GET_PRODUCTS = `${BASE_URL}/products`;
  const URL_DELETE_PRODUCT = `${BASE_URL}/product`;

  const [products, setProducts] = useState<Product[]>([]);
  const [productID, setProductID] = useState<string>("");
  const [empty, setEmpty] = useState<boolean>(false);

  const fetchAllProduct = async () => {
    try {
      // Fetch all products from the backend
      const response = await axios.get(URL_GET_PRODUCTS);

      // Check if the response contains data
      if (response.data.result !== null) {
        // Set the products state with the fetched data
        setProducts(response.data.result);
        setEmpty(false);
      } else {
        setEmpty(true);
      }
    } catch (error) {
      // show error message using Swal
      Swal.fire({
        icon: "error",
        title: "เกิดข้อผิดพลาด",
        text: "ไม่สามารถดึงข้อมูลสินค้าทั้งหมดได้",
      });
      console.error("Error fetching products:", error);
    }
  };

  useEffect(() => {
    // Fetch all products when the component mounts
    fetchAllProduct();
  }, []);

  const validateProductID = (productId: string) => {
    // Check if product ID matches the pattern: xxxxx-xxxxx-xxxxx-xxxxx-xxxxx-xxxxx
    const productIdPattern =
      /^[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}-[A-Z0-9]{5}$/;

    if (!productIdPattern.test(productId)) {
      return false;
    }

    // // Additional check to ensure it contains at least one number and one uppercase letter
    // const hasNumber = /\d/.test(productId);
    // const hasUppercase = /[A-Z]/.test(productId);

    // if (!hasNumber || !hasUppercase) {
    //   return false;
    // }

    // console.log("Product ID:", productId);

    return true;
  };

  const handleProductIDChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    // Convert input to uppercase
    const inputValue = e.target.value.toUpperCase();
    // Set the product ID state
    setProductID(inputValue);
    // Add a hyphen after every 5 characters
    if ((inputValue.length + 1) % 6 === 0 && inputValue.length < 35) {
      setProductID((prev) => prev + "-");
    }
  };

  const addProduct = async () => {
    try {
      // Validate the product ID before sending it to the backend
      if (!validateProductID(productID)) {
        Swal.fire({
          icon: "error",
          title: "รหัสสินค้าไม่ถูกต้อง",
          text: "รหัสสินค้าต้องมีความยาว 30 ตัวอักษร เป็นตัวอักษรภาษาอังกฤษพิมพ์ใหญ่หรือตัวเลขเท่านั้น",
        });
        return;
      }

      // Add the product to the backend
      await axios.post(URL_ADD_PRODUCT, { product_id: productID });

      // Show success message
      Swal.fire({
        icon: "success",
        title: "เพิ่มสินค้าสำเร็จ",
        text: `สินค้ารหัส ${productID} ถูกเพิ่มเรียบร้อยแล้ว`,
      });

      // Refresh the product list after adding a new product
      fetchAllProduct();
    } catch (error: any) {
      console.error("Error adding product:", error);
      Swal.fire({
        icon: "error",
        title: "เกิดข้อผิดพลาด",
        text: error.response?.data?.error || "ไม่สามารถเพิ่มสินค้าได้",
      });
    }
  };

  const deleteProduct = async (productId: string) => {
    try {
      await axios.delete(URL_DELETE_PRODUCT, { params: { id: productId } });
      Swal.fire({
        icon: "success",
        title: "ลบข้อมูลสำเร็จ",
        text: `สินค้ารหัส ${productId} ถูกลบเรียบร้อยแล้ว`,
      });
    } catch (error: any) {
      console.error("Error deleting product:", error);
      Swal.fire({
        icon: "error",
        title: "เกิดข้อผิดพลาด",
        text: error.response?.data?.error || "ไม่สามารถลบสินค้าได้",
      });
    }
  };

  const showConfirmDeleteModal = (productId: string) => {
    Swal.fire({
      title: " ต้องการลบข้อมูลใช่ไหม?",
      text: `ต้องการลบสินค้านี้ ${productId} ใช่ไหม?`,
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "ลบข้อมูล",
      cancelButtonText: "ยกเลิก",
    }).then(async (result) => {
      if (result.isConfirmed) {
        await deleteProduct(productId);
        fetchAllProduct(); // Refresh the product list after deletion
      }
    });
  };

  const showQRCodeModal = () => {
    Swal.fire({
      imageUrl: qrcodeImage,
      imageAlt: "A tall image",
      width: 400,
      padding: "3em",
    });
  };

  const RenderProducts = () => {
    if (empty) {
      return (
        <div className="text-center text-gray-500">
          <p>ไม่มีสินค้าที่เพิ่มไว้</p>
          <p>กรุณาเพิ่มสินค้าก่อน</p>
        </div>
      );
    } else {
      return (
        <div className="flex items-center gap-4">
          <table className="min-w-full bg-white border border-gray-300">
            <thead>
              <tr>
                <th className="px-4 py-2 border-b">ID</th>
                <th className="px-4 py-2 border-b">Product ID</th>
                <th className="px-4 py-2 border-b">View</th>
                <th className="px-4 py-2 border-b">Delete</th>
              </tr>
            </thead>
            <tbody>
              {products.map((product: Product, index: number) => (
                <tr key={product.product_id}>
                  <td className="px-4 py-2 border-b">{index + 1}</td>
                  <td className="px-4 py-2 border-b">{product.product_id}</td>
                  <td className="px-4 py-2 border-b">
                    <button
                      onClick={showQRCodeModal}
                      className="px-3 py-1 bg-green-400 text-white rounded hover:bg-green-500"
                    >
                      QR
                    </button>
                  </td>
                  <td className="px-4 py-2 border-b">
                    <button
                      onClick={() => showConfirmDeleteModal(product.product_id)}
                      className="px-3 py-1 bg-red-500 text-white rounded hover:bg-red-600"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      );
    }
  };

  return (
    <>
      <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
        <h1 className="text-2xl font-bold mb-6">ระบบ Product Management</h1>
        <div className="flex flex-row justify-center items-center mb-4">
          <div>รหัสสินค้า</div>
          <input
            type="text"
            className="ml-2 px-3 py-2 border border-gray-300 rounded uppercase"
            placeholder="xxxxx-xxxxx-xxxxx-xxxxx-xxxxx-xxxxx"
            value={productID}
            onChange={handleProductIDChange}
            style={{ width: "500px" }}
            maxLength={35}
          />
          <button
            onClick={addProduct}
            className="ml-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
          >
            Add
          </button>
        </div>
        <RenderProducts />
      </div>
    </>
  );
}

export default App;
