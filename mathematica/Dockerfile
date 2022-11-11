FROM btwiuse/mathematica:sh as mathematica

FROM btwiuse/arch

COPY --from=mathematica /*.sh /mathematica.sh

RUN bash /mathematica.sh -- -auto -verbose
