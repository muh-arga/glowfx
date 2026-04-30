import { useEffect, useState } from "react";
import { createProduct, updateProduct } from "../../api/productApi";

export default function ProductFormModal({
  isOpen,
  onClose,
  onSuccess,
  initialData,
}) {
  const isEdit = !!initialData;

  const defaultForm = {
    name: "",
    sku: "",
    status: "active",
  };

  const [form, setForm] = useState(defaultForm);
  const [errors, setErrors] = useState({});
  const [loading, setLoading] = useState(false);

  const resetForm = () => {
    setForm(defaultForm);
    setErrors({});
  };

  useEffect(() => {
    if (isOpen) {
      if (initialData) {
        setForm({
          name: initialData.name || "",
          sku: initialData.sku || "",
          status: initialData.status || "active",
        });
      } else {
        resetForm();
      }
    }
  }, [initialData, isOpen]);

  const handleClose = () => {
    resetForm();
    onClose();
  };

  const handleSubmit = async () => {
    setLoading(true);
    setErrors({});

    try {
      if (isEdit) {
        await updateProduct(initialData.id, form);
      } else {
        await createProduct(form);
      }

      onSuccess();
      resetForm();
      onClose();
    } catch (err) {
      const res = err.response?.data;

      if (res?.errors) {
        setErrors(res.errors);
        return;
      }

      alert(res?.message || "something went wrong");
    } finally {
      setLoading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div style={styles.overlay}>
      <div style={styles.modal}>
        <h3>{isEdit ? "Edit Product" : "Create Product"}</h3>

        <div>
          <input
            placeholder="Name"
            value={form.name}
            onChange={(e) => {
              setForm({ ...form, name: e.target.value });
              setErrors({ ...errors, name: null });
            }}
            style={{
              border: errors.name ? "1px solid red" : "1px solid #ddd",
            }}
          />
          {errors.name && <p style={styles.error}>{errors.name}</p>}
        </div>

        <div>
          <input
            placeholder="SKU"
            value={form.sku}
            onChange={(e) => {
              setForm({ ...form, sku: e.target.value });
              setErrors({ ...errors, sku: null });
            }}
            style={{
              border: errors.sku ? "1px solid red" : "1px solid #ddd",
            }}
          />
          {errors.sku && <p style={styles.error}>{errors.sku}</p>}
        </div>

        <div>
          <select
            value={form.status}
            onChange={(e) =>
              setForm({ ...form, status: e.target.value })
            }
          >
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
          </select>
        </div>

        <div style={{ marginTop: 10 }}>
          <button onClick={handleClose}>Cancel</button>
          <button onClick={handleSubmit} disabled={loading}>
            {loading ? "Saving..." : "Save"}
          </button>
        </div>
      </div>
    </div>
  );
}

const styles = {
  overlay: {
    position: "fixed",
    inset: 0,
    background: "rgba(0,0,0,0.4)",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  modal: {
    background: "#fff",
    padding: 20,
    borderRadius: 10,
    width: 320,
  },
  error: {
    color: "red",
    fontSize: 12,
  },
};