data "aws_route53_zone" "host_domain" {
  name = local.host_domain
}

resource "aws_route53_record" "cert_validation" {
  for_each = {
    for dvo in aws_acm_certificate.cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
      type   = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = data.aws_route53_zone.host_domain.zone_id
}

resource "aws_acm_certificate" "cert" {
  domain_name       = local.host_domain
  validation_method = "DNS"

  tags = {
    Environment = var.app_environment
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}

resource "aws_route53_record" "site" {
  zone_id = data.aws_route53_zone.host_domain.zone_id
  name    = local.host_domain
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.static-skhole.domain_name
    zone_id                = aws_cloudfront_distribution.static-skhole.hosted_zone_id
    evaluate_target_health = false
  }
}