variable "aws_region" {
  description = "Región de AWS para el despliegue"
  type        = string
  default     = "us-east-1"
}

variable "db_password" {
  description = "Contraseña de la base de datos PostgreSQL en RDS"
  type        = string
  sensitive   = true
}